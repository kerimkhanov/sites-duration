FROM golang:1.18.6

WORKDIR /monitoring

# Copy the Pre-built binary job
COPY build/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
