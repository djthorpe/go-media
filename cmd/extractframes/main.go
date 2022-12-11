/*
An example of extracting frames from a video. Provide a single video file on the command line,
and it will extract and save individual frames to the current working directory, or the directory
specified by the -out flag.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	// Packages
	media "github.com/mutablelogic/go-media/pkg/media"

	// Namespace imports
	. "github.com/mutablelogic/go-media"
)

var (
	flagDebug = flag.Bool("debug", false, "Enable debug output")
	flagOut   = flag.String("out", "", "Output directory for artwork")
)

func main() {
	// Check arguments
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(-1)
	}

	// Create a cancellable context
	ctx := ContextForSignal(os.Interrupt)

	// Create a media manager, set debugging
	manager := media.New()
	manager.SetDebug(*flagDebug)

	// Check the out directory
	if *flagOut != "" {
		if info, err := os.Stat(*flagOut); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-2)
		} else if !info.IsDir() {
			fmt.Fprintf(os.Stderr, "%q: not a directory\n", info.Name())
			os.Exit(-2)
		}
	}

	// Open the file
	media, err := manager.OpenFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-2)
	}

	// Find the video stream for extraction
	streams := media.StreamsByType(MEDIA_FLAG_VIDEO)
	if len(streams) != 1 {
		fmt.Fprintln(os.Stderr, "No video stream found, or more than one video stream found")
		os.Exit(-2)
	}

	// Decode the media
	if err := manager.Decode(ctx, media, func(_ context.Context, packet Packet) error {
		if streams[0] == packet.Stream() {
			fmt.Println("PACKET", packet)
		}
		return nil
	}); err != nil && err != context.Canceled {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-2)
	}

	// If the context was cancelled, print a message
	if ctx.Err() != nil {
		fmt.Fprintln(os.Stderr, "\nInterrupted")
	}
}

// ContextForSignal returns a context object which is cancelled when a signal
// is received. It returns nil if no signal parameter is provided
func ContextForSignal(signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		return nil
	}

	ch := make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())

	// Send message on channel when signal received
	signal.Notify(ch, signals...)

	// When any signal received, call cancel
	go func() {
		<-ch
		cancel()
	}()

	// Return success
	return ctx
}
