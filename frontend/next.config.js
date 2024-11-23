/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    onDemandEntries: {
      maxInactiveAge: 25 * 1000,
      pagesBufferLength: 2,
    },
    experimental: {
      serverActions: {
        bodySizeLimit: '2mb'
      },
    },
    async rewrites() {
      return [
        {
          source: '/api/:path*',
          destination: 'http://backend:8000/:path*',
        },
      ]
    },
    webpack: (config) => {
      config.watchOptions = {
        poll: 1000,
        aggregateTimeout: 300,
      }
      return config
    },
}

module.exports = nextConfig