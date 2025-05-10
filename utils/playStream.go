package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// a function that checks if the channel is live and returns the stream and pipes it to the player yt-dlp and vlc are needed

//uses yt-dlp and vlc
func PlayStream(channelName string) error {
	// Run yt-dlp and capture output
	ytDlpCmd := exec.Command("yt-dlp", "--print", "urls", "-q", fmt.Sprintf("twitch.tv/%s", channelName))
	output, err := ytDlpCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run yt-dlp: %w", err)
	}

	streamURL := strings.TrimSpace(string(output))
	if streamURL == "" {
		return fmt.Errorf("no stream URL found for channel: %s", channelName)
	}

	// Run VLC with the stream URL
	vlcCmd := exec.Command("vlc", "--play-and-exit", "--network-caching=100", streamURL)
	err = vlcCmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start VLC: %w", err)
	}

	return nil
}
// uses streamlinks 


func _PlayTwitchChannel(channel string) error {
	url := fmt.Sprintf("https://www.twitch.tv/%s", channel)

	cmd := exec.Command(
		"streamlink",
		"--twitch-low-latency",
		"--hls-live-edge=1",
		"--quiet",
		"--player=vlc",
		url,
		"best",
	)

	return cmd.Run()
}


var twitchProcess *os.Process

func PlayTwitchChannel(channel string) error {
   var player string
   
   if IsAppInstalled("mpv") {
       player = "mpv"
   } else {
       player = "vlc"
       
   }
    url := fmt.Sprintf("https://www.twitch.tv/%s", channel)

    cmd := exec.Command(
        "streamlink",
        "--twitch-low-latency",
        "--hls-live-edge=1",
        "--quiet",
        "--player="+player,
        url,
        "best",
    )

    // Set process group for Unix-like systems
    //!! can't be opened on the windows version of GO
    if runtime.GOOS != "windows" {
        cmd.SysProcAttr = &syscall.SysProcAttr{
            Setpgid: true,
        }
    }

    err := cmd.Start()
    if err != nil {
        return err
    }

    // Save process reference for later termination
    twitchProcess = cmd.Process
    //fmt.Printf("Twitch stream started with PID %d\n", twitchProcess.Pid)
    return nil
}

func StopTwitchChannel() error {
    if twitchProcess == nil {
        return fmt.Errorf("no Twitch stream process to stop")
    }

    var err error
    if runtime.GOOS == "windows" {
        err = twitchProcess.Kill()
    } else {
        // Kill the process group on Unix-like systems
        //!! can't be opened on the windows version of GO
        pgid, pgErr := syscall.Getpgid(twitchProcess.Pid)
        if pgErr != nil {
            return fmt.Errorf("failed to get pgid: %w", pgErr)
        }
        err = syscall.Kill(-pgid, syscall.SIGKILL)
    }

    if err != nil {
        return fmt.Errorf("failed to kill process: %w", err)
    }
    twitchProcess = nil
    return nil
}

func IsAppInstalled(appName string) bool {
	// First attempt: Check if the command exists in PATH
	_, err := exec.LookPath(appName)
    return err == nil
}