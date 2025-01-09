# Job Posting Application

## Requirements

- Go
- Air
- Buf
- Docker
- Proto

## Steps to Run the Application Locally

1. Install the required packages:

   ```sh
   go mod tidy
   ```

2. Initialize Air:

   ```sh
   air init
   ```

3. Run the gateway:

   ```sh
   go run main.go
   ```

   or if using Air:

   ```sh
   air
   ```

4. Run the auth service:

   - Open a new shell
   - Navigate to the `auth` folder
   - Run the service:
     ```sh
     go run main.go
     ```

5. Run the company service:

   - Open a new shell
   - Navigate to the `company` folder
   - Run the service:
     ```sh
     go run main.go
     ```

6. Run the job service:
   - Open a new shell
   - Navigate to the `job` folder
   - Run the service:
     ```sh
     go run main.go
     ```

## Installing Buf

To install Buf, follow these steps:

1. Download the Buf binary from the [official Buf website](https://docs.buf.build/installation).

2. Move the binary to a directory included in your system's `PATH`. For example:

   ```sh
   mv buf /usr/local/bin/
   ```

3. Verify the installation by running:

   ```sh
   buf --version
   ```

   You should see the version of Buf that you installed.
