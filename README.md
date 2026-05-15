# Kazumi 
## Work in Progress

Kazumi is a command-line tool designed to transform tracks into high-quality Nightcore. It uses yt-dlp for downloads and FFmpeg for audio manipulation, offering a streamlined workflow for music processing.

---

## Features

- **YouTube Integration**: Download audio directly from URLs with automatic metadata extraction.
- **Nightcore Processing**:
    - **Speed & Tempo**: Fine-tuned speed control (0.5x - 2.0x).
    - **Pitch Shifting**: Authentic Nightcore pitch adjustment.
    - **Reverb & Echo**: Customizable reverb parameters for spatial effects.
- **Workspace Management**: Automatic management of caches, logs, and output folders.
- **Dependency Checks**: Built-in verification to ensure required tools are installed.

---

## Prerequisites

Kazumi requires the following external tools to be present in your system's PATH:

1.  **Go** (v1.18+) - To compile the source.
2.  **yt-dlp** - For audio downloading.
3.  **FFmpeg** - For audio processing.

---

## Installation and Building

### Linux
1.  **Install Dependencies**:
    ```bash
    sudo apt update && sudo apt install yt-dlp ffmpeg  # Debian/Ubuntu
    sudo dnf install yt-dlp ffmpeg                    # Fedora
    sudo pacman -S yt-dlp ffmpeg                      # Arch Linux
    ```
2.  **Build**:
    ```bash
    make build
    # Or manually:
    go build -o bin/kazumi cmd/kazumi/main.go
    ```

### macOS
1.  **Install Dependencies** (via Homebrew):
    ```bash
    brew install yt-dlp ffmpeg
    ```
2.  **Build**:
    ```bash
    make build
    # Or manually:
    go build -o bin/kazumi cmd/kazumi/main.go
    ```

### Windows
1.  **Install Dependencies**:
    - Download `yt-dlp.exe` from the official GitHub releases and add it to your PATH.
    - Download `ffmpeg` from Gyan.dev and add the `bin` folder to your PATH.
2.  **Build**:
    ```powershell
    # Using PowerShell
    go build -o bin/kazumi.exe cmd/kazumi/main.go
    ```

---

## Usage

Once built, you can run Kazumi from the bin directory. The basic syntax is:

```bash
./bin/kazumi [flags] <audio_url>
```

### Flags

- `-s, --speed [value]`: Adjust the speed of the output. If the flag is provided without a value, it defaults to `1.25`.
- `-p, --pitch [value]`: Adjust the pitch of the output. If the flag is provided without a value, it defaults to `1.33`.
- `-r, --reverb`: Enable reverb effect for the audio. 
- `-i, --image <url>`: Specify an image URL to be used for the video background. Default: YT Video's Thumbnail
- `-h, --help`: Display the help message.
- `-j, --json`: Use JSON file, to download in bulk
- `-a, --audio`: Only download audio.

### Examples

**Basic Nightcore (default speed/pitch boost):**
```bash
./bin/kazumi --speed --pitch https://www.youtube.com/watch?v=example

./bin/kazumi -s -p https://www.youtube.com/watch?v=example
```


**Custom Speed and Pitch with Reverb:**
```bash
./bin/kazumi -s=1.3 -p=1.4 -r https://www.youtube.com/watch?v=example
```

**With a custom image:**
```bash
./bin/kazumi -i "https://example.com/image.jpg" https://www.youtube.com/watch?v=example
```


**Use JSON:**
```bash
./bin/kazumi --json /home/user/data.json
```

---

## Project Structure

```text
├── cmd/
│   ├── kazumi/    # Main entry point
│   └── cli/       # CLI command definitions (Cobra) [WIP]
├── lib/
│   ├── audio/     # Processing & Downloading logic
│   ├── image/     # Video generation
│   └── misc/      # Utils, Logging, Dependencies
└── ui/            # Planned UI implementations (TUI/GUI/Web)
```

---

## Roadmap / TODO

- [x] Implement command line functionality
- [x] Generate a simple video using the audio and image provided
- [x] Allow multiple videos to be batch processed (JSON config).
- [x] Fix the weird audio issues
- [ ] Add a config to store information such as locations
- [ ] Add a history database
- [ ] Add more audio filters like Bass Boost or EQ presets.
- [ ] Add verbose error messages.
- [ ] Add a way to automatically download & cache the dependencies if not present.
- [ ] Allow you to set custom value for reverb in CLI
- [ ] Allow user to use local audio

### UI Implementations
- [ ] Build a terminal user interface using Bubble Tea or similar.
- [ ] Use Fyne to build a GUI Application
- [ ] Create a basic website to manage the download remotely

---

## Contributing

Contributions are welcome. Feel free to:
- Report bugs by opening an issue.
- Suggest new features.
- Submit pull requests for improvements.

---

## License

Distributed under the MIT License. See `LICENSE` for more information.
