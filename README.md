# Video Editing Program

This is a simple command-line program written in Go for video editing. It uses the FFmpeg library to edit video files by trimming them based on user input.

## Prerequisites

Before using this program, make sure you have the following prerequisites installed on your system:

- Go (https://golang.org/doc/install)
- FFmpeg (https://www.ffmpeg.org/download.html)

## Installation

1. Clone this repository to your local machine:

   ```shell
   git clone <repository-url>
   ```

2. Navigate to the project directory:

   ```shell
   cd <repository-directory>
   ```

3. Build the program:

   ```shell
   go build videoEditing.go
   ```

## Usage

To use this program, follow these steps:

1. Open a terminal window.

2. Navigate to the directory where the program binary is located.

3. Run the program with the following command:

   ```shell
   ./videoEditing <integer>
   ```

   Replace `<integer>` with the number of seconds you want to trim from the video's end.

4. The program will process the video located at "./testdata/in1.mp4" and trim it according to your input.

5. The trimmed video will be saved to "./outdir/out2.mp4."

## Example

Here's an example of how to use the program:

```shell
./videoEditing 10
```

This command will trim 10 seconds from the end of the input video and save the result as "./outdir/out2.mp4."

## Note

- Make sure that the input video file is located at "./testdata/in1.mp4" and that the output directory "./outdir" exists before running the program.

## License

This program is open-source and is distributed under the [MIT License](LICENSE).

## Author

This program was created by eduwr.

If you have any questions or encounter any issues, please feel free to contact me