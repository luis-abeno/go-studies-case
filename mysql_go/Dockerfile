# Use the official image as a parent image
FROM mysql:8.0

# Set the working directory in the container
WORKDIR /docker-entrypoint-initdb.d

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=testing_go_mysql
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=root

# Expose port 3306
EXPOSE 3306

# Run the command on container startup
CMD ["mysqld"]