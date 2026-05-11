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

Once built, you can run Kazumi from the bin directory:

```bash
./bin/kazumi
```

*Note: The CLI is currently under active development and hasn't been properly implemented yet.*

---

## Project Structure

```text
├── cmd/
│   ├── kazumi/    # Main entry point
│   └── cli/       # CLI command definitions (Cobra) [WIP]
├── lib/
│   ├── audio/     # Processing & Downloading logic
│   ├── image/     # Video generation [WIP]
│   └── misc/      # Utils, Logging, Dependencies
└── ui/            # Planned UI implementations (TUI/GUI/Web)
```

---

## Roadmap / TODO

- [ ] Implement command line functionality
- [ ] Allow multiple videos to be batch processed.
- [ ] Add a config to store information such as locations
- [ ] Generate a simple video using the audio and image provided
- [ ] Add more audio filters like Bass Boost or EQ presets.
- [ ] Add verbose error messages.
- [ ] Add a way to automatically download & cache the dependencies if not present.

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
