package main

import (
	"bufio"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	scanner bufio.Scanner
	// client appproto.ProductServiceClient // client implementation of grpc proto, update the name accordingly
}

func NewProductClient(
	reader io.Reader,
	// preferably add interceptor class as param here, but can also be added directly
) ProductClient {
	// initialize grpc and add interceptor here
	conn, err := grpc.Dial(":9000",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // to allow insecure connection
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	log.Println("conn:", conn) // just to avoid warning

	return ProductClient{
		scanner: *bufio.NewScanner(reader),
		// also save proto client to be used later
	}
}

func (pc *ProductClient) MainLoop() {

	for {
		log.Println("Simple Client Command")
		log.Println("---------------------")
		log.Println("1) Add Product")
		log.Println("2) Get All Product")
		log.Println("3) Get Product by Id")
		log.Print("->: ")
		// Scans a line from Stdin(Console)
		pc.scanner.Scan()
		// Holds the string that scanned
		text := pc.scanner.Text()

		switch text {
		case "1":
			pc.HandleAddProduct()
		case "2":
			pc.HandleGetProducts()
		case "3":
			pc.HandleGetProduct()
		}
		log.Println("---------------------")
	}
}

func (pc *ProductClient) HandleAddProduct() {
	log.Println("Input product data")
	log.Println("---------------------")
	log.Println("id:")
	pc.scanner.Scan()
	productID := pc.scanner.Text()
	log.Println("name:")
	pc.scanner.Scan()
	productName := pc.scanner.Text()
	log.Println("description:")
	pc.scanner.Scan()
	productDesc := pc.scanner.Text()

	// Send this data to grpc server
	log.Println("inputted data product. id:", productID, ", name:", productName, ", desc:", productDesc)
}

func (pc *ProductClient) HandleGetProducts() {
	// hit grpc server and print all response
}

func (pc *ProductClient) HandleGetProduct() {
	log.Println("Input product data")
	log.Println("---------------------")
	log.Println("id:")
	pc.scanner.Scan()
	productID := pc.scanner.Text()
	// Send this data to grpc server
	log.Println("inputted data product. id:", productID)

	// also print the response
}
