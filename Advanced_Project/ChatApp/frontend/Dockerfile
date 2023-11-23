# Stage 1: Build stage
FROM node:21-alpine AS builder

WORKDIR /app

COPY ./frontend/package.json .

# Install dependencies for building
RUN npm install 

COPY ./frontend .

# Build the application
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
