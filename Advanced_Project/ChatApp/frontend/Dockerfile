# Stage 1: Build stage
FROM node:21-alpine AS builder

WORKDIR /app

COPY package.json .

# Install dependencies for building
RUN npm install --production 

COPY . .

# Build the application
RUN npm run build

# Stage 2: Final stage
FROM node:21-alpine

WORKDIR /app

# Copy only necessary files from the builder stage
COPY --from=builder /app/package.json .
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist

EXPOSE 3000

CMD ["npm", "start"]
