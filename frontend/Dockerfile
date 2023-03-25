# Base image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package.json and yarn.lock to container
COPY package*.json yarn.lock ./

# Install dependencies with yarn
RUN yarn install --production=false

# Copy the rest of the application files to container
COPY . .

# Build the application
RUN yarn build

# Expose port 3000 for the application
EXPOSE 3001

# Start the application
CMD ["yarn", "start"]