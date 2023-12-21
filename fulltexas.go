package fulltexas

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Options struct {
	LogFormat string
	LogLevel  string
}

func Execute() int {
	options := parseArgs()

	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	err = run(options)
	if err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}
	return 0
}

func parseArgs() Options {
	const logLevelHelp = "Log level (debug, info, warn, error), default: info"

	options := Options{}

	ll := flag.String("ll", "info", fmt.Sprintf("%s (shorthand)", logLevelHelp))
	flag.StringVar(&options.LogLevel, "log-level", *ll, logLevelHelp)

	flag.StringVar(&options.LogFormat, "log-format", "text", "Log format (text or json)")

	flag.Parse()

	return options
}

func run(options Options) error {
	slog.Debug("test", "test", "Debug")
	slog.Debug("test", "LogLevel", options.LogLevel)
	slog.Info("test", "test", "Info")
	slog.Error("test", "test", "Error")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	imgData1, err := os.ReadFile("images/turtle1.png")
	if err != nil {
		log.Fatal(err)
	}

	imgData2, err := os.ReadFile("images/turtle2.png")
	if err != nil {
		log.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("png", imgData1),
		genai.ImageData("png", imgData2),
		genai.Text("Describe the difference between these two pictures, with scientific detail"),
	}
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(bs))
	return nil
}
