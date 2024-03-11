package main

import (
	pb "DistributedFileSystem/dfs" // Import the generated Go code
	"context"
	"fmt"
	"io"
	"net"
	"os"

	// "github.com/abema/go-mp4"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDFSServer
	client pb.DFSClient
}

func (s *server) NotifyClient(ctx context.Context, req *pb.NotifyClientRequest) (*pb.Empty, error) {
	fmt.Println("NotifyClient called:" + req.GetMessage())
	return &pb.Empty{}, nil
}
func main() {
	// master port
	masterPort := "8080"
	// client port
	clientPort := ":8081"

	// connect to the master tracker
	masterConn, err := grpc.Dial("localhost:"+masterPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect to master:", err)
		return
	}
	defer masterConn.Close()
	cMaster := pb.NewDFSClient(masterConn)

	// create a listener on the client port waiting for teh success or failure of the operation after the master tracker has finished the operation
	lis, err := net.Listen("tcp", clientPort)
	if err != nil {
		fmt.Println("failed to listen on the master:", err)
		return
	}
	defer lis.Close()
	s := grpc.NewServer()
	srv := &server{
		client: cMaster,
	}
	pb.RegisterDFSServer(s, srv)

	// Read input from user to know if it's a file upload or a file download
	fmt.Print("Choose an option (1 for file upload, 2 for file download): ")
	var text string
	fmt.Scanln(&text)

	if text == "1" {
		//1. send to the master tracker request to upload
		resToUpload, err := cMaster.RequestToUpload(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Println("Error calling RequestToUpload:", err)
			return
		}
		//2. extract the data keeper port number from the response
		dataNodePort := resToUpload.GetToken()
		fmt.Println("Data Keeper Port Numebr :", dataNodePort)

		//3. connect to the data keeper
		dataConn, err := grpc.Dial("localhost:"+dataNodePort, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect to data keeper:", err)
			return
		}
		defer dataConn.Close()
		cData := pb.NewDFSClient(dataConn)

		// Read input from user
		fmt.Print("Enter the file path: ")
		var filePath string
		fmt.Scanln(&filePath)
		fmt.Print("Enter the file name: ")
		var fileName string
		fmt.Scanln(&fileName)

		// Read the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		defer file.Close()
		// Read the MP4 file as bytes
		mp4Bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading MP4 file:", err)
			return
		}
		//4. send request to the data keeper to upload the file
		resToUploadFile, err := cData.UploadFile(context.Background(), &pb.UploadFileRequest{FileName: fileName, FileData: mp4Bytes})
		if err != nil {
			fmt.Println("Error calling UploadFile:", err)
			return
		}
		fmt.Println("Data Keeper response:", resToUploadFile)

	} else if text == "2" {
		// Read input from user
		fmt.Print("Enter the file name: ")
		var fileName string
		fmt.Scanln(&fileName)
		//1. send to the master tracker request to download
		resToDownload, err := cMaster.RequestToDownload(context.Background(), &pb.RequestToDownloadRequest{FileName: fileName})
		if err != nil {
			fmt.Println("Error calling RequestToDownload:", err)
			return
		}
		//2. extract the data keeper port number from the response
		dataNodePort := string(resToDownload.GetMachineInfos()[0].GetPort()) // get the first data keeper
		fmt.Println("Data Keeper Port Numebr :", dataNodePort)

		//3. connect to the data keeper
		dataConn, err := grpc.Dial("localhost:"+dataNodePort, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect to data keeper:", err)
			return
		}
		defer dataConn.Close()
		cData := pb.NewDFSClient(dataConn)

		//4. send request to the data keeper to download the file
		resToDownloadFile, err := cData.DownloadFile(context.Background(), &pb.DownloadFileRequest{FileName: fileName})
		if err != nil {
			fmt.Println("Error calling DownloadFile:", err)
			return
		}
		// fmt.Println("Data Keeper response:", resToDownloadFile)

		//5. save the file to the local file system
		err = os.WriteFile(fileName, resToDownloadFile.FileData, 0644)
		if err != nil {
			fmt.Println("Failed to save file:", err)
			return
		}
		fmt.Println("File saved successfully:", fileName)

	} else {
		fmt.Println("Invalid option")
	}

	// 6. wait for the master response to know the result of the operation
	fmt.Printf("Client started Listening on port %s ...\n", clientPort)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
