# syntax=docker/dockerfile:1

# Build arguments
ARG NODE_VERSION=20
ARG PNPM_VERSION=9.15.4

# Base stage - common for all environments
FROM node:${NODE_VERSION}-alpine AS base

# Install pnpm
RUN corepack enable && corepack prepare pnpm@${PNPM_VERSION} --activate

WORKDIR /app

# Install dependencies only when needed
FROM base AS deps

# Copy package files
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

# Install dependencies based on the lock file
RUN pnpm install --frozen-lockfile

# Development stage
FROM base AS development

# Copy node_modules from deps stage
COPY --from=deps /app/node_modules ./node_modules

# Copy source code
COPY . .

# Environment variables for development
ENV NODE_ENV=development \
    NEXT_TELEMETRY_DISABLED=1

# Expose port
EXPOSE 3000

# Run development server
CMD ["pnpm", "dev"]

# Builder stage for production
FROM base AS builder

# Copy node_modules from deps stage
COPY --from=deps /app/node_modules ./node_modules

# Copy source code
COPY . .

# Environment variables for building
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1

# Build the application
RUN pnpm build

# Production stage
FROM node:${NODE_VERSION}-alpine AS production

WORKDIR /app

# Install dumb-init for proper signal handling
RUN apk add --no-cache dumb-init

# Create non-root user
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

# Copy necessary files from builder
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

# Set ownership to non-root user
RUN chown -R nextjs:nodejs /app

# Use non-root user
USER nextjs

# Environment variables for production
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000 \
    HOSTNAME=0.0.0.0

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
    CMD node -e "require('http').get('http://localhost:3000/api/health', (r) => {process.exit(r.statusCode === 200 ? 0 : 1)})"

# Run the application with dumb-init
ENTRYPOINT ["dumb-init", "--"]
CMD ["node", "server.js"]
