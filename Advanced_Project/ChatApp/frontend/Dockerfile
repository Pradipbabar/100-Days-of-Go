# Stage 1: Build stage
FROM node:21-alpine AS builder

WORKDIR /app

COPY package.json .

# Install dependencies for building
RUN npm install 

COPY . .

# Build the application
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
