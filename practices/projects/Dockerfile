# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Copy the 'page' folder and its files, such as 'about.html' and 'contact.html'
COPY page /app/page

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go source code and the page directory into the container
COPY . .

# Build the Go app using the COMPILE_FILE environment variable
RUN go build -o webpage main.go

# Expose the port the app will run on
EXPOSE 10000

# Command to run the application
CMD ["./webpage"]
