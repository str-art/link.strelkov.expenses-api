/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  env: {
    API_URL: process.env.API_URL,
  },
  images: {
    unoptimized: true,
  },
  output: "standalone",
  swcMinify: true,
};

module.exports = nextConfig;
