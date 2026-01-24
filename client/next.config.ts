import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Enable standalone output for Docker production builds
  output: "standalone",
  
  // Disable x-powered-by header for security
  poweredByHeader: false,
  
  // Enable compression
  compress: true,
};

export default nextConfig;
