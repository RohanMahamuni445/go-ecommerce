# Step 1: Build the React app
FROM node:18 AS build

WORKDIR /app

# Copy package.json and lock file first for caching
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app
COPY . .

# Build the app for production
RUN npm run build

# Step 2: Serve the build using Nginx
FROM nginx:alpine

# Copy the build folder from previous stage to nginx html directory
COPY --from=build /app/build /usr/share/nginx/html

# Remove default nginx config and replace with our custom one if needed
# (Optional: if you have a custom nginx.conf, otherwise this line is not needed)
# COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80 to outside
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]

