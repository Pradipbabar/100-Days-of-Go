# Stage 1: Build stage
FROM node:14-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache --virtual .build-deps \
    python \
    make \
    g++ \
    && apk del .build-deps

COPY package.json .

# Install dependencies for building
RUN npm install --production --force

COPY . .

# Build the application
RUN npm run build

# Stage 2: Final stage
FROM node:14-alpine

WORKDIR /app

# Copy only necessary files from the builder stage
COPY --from=builder /app/package.json .
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist

EXPOSE 3000

CMD ["npm", "start"]
