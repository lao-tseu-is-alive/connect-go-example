//go:build integration

package main

import (
	"context"
	"net/http"
	"os"
	"testing"

	"connectrpc.com/connect"
	greetv1 "github.com/lao-tseu-is-alive/connect-go-example/gen/greet/v1"
	"github.com/lao-tseu-is-alive/connect-go-example/gen/greet/v1/greetv1connect"
)

// Integration benchmarks require the server to be running:
//   cd cmd/server && LOG_LEVEL=error go run server.go
//
// Run with:
//   go test -tags=integration -bench=. -benchmem ./cmd/client/...

func getServerURL() string {
	if url := os.Getenv("SERVER_URL"); url != "" {
		return url
	}
	return "http://127.0.0.1:8080"
}

// BenchmarkIntegration_Connect benchmarks real server with Connect protocol
func BenchmarkIntegration_Connect(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup: make a few requests to establish connection
	for i := 0; i < 10; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Skipf("Server not available: %v", err)
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// BenchmarkIntegration_GRPC benchmarks real server with gRPC protocol
func BenchmarkIntegration_GRPC(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
		connect.WithGRPC(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup
	for i := 0; i < 10; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Skipf("Server not available: %v", err)
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// BenchmarkIntegration_JSON benchmarks real server with JSON encoding
func BenchmarkIntegration_JSON(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
		connect.WithProtoJSON(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup
	for i := 0; i < 10; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Skipf("Server not available: %v", err)
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// --- Parallel Integration Benchmarks ---

// BenchmarkIntegration_Connect_Parallel benchmarks real server with concurrent requests
func BenchmarkIntegration_Connect_Parallel(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup
	_, err := client.Greet(context.Background(), req)
	if err != nil {
		b.Skipf("Server not available: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}

// BenchmarkIntegration_GRPC_Parallel benchmarks real server with gRPC concurrent requests
func BenchmarkIntegration_GRPC_Parallel(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
		connect.WithGRPC(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup
	_, err := client.Greet(context.Background(), req)
	if err != nil {
		b.Skipf("Server not available: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}

// BenchmarkIntegration_JSON_Parallel benchmarks real server with JSON concurrent requests
func BenchmarkIntegration_JSON_Parallel(b *testing.B) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		getServerURL(),
		connect.WithProtoJSON(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Warmup
	_, err := client.Greet(context.Background(), req)
	if err != nil {
		b.Skipf("Server not available: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}
