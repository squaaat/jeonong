// next.config.js에서 promise, async function 안먹음.
// 거지같은 next.config.js
module.exports = {
  env: {
    J_ENV: process.env.J_ENV,
    J_CICD: process.env.J_CICD,
    GOOGLE_OAUTH_CLIENT_ID: process.env.GOOGLE_OAUTH_CLIENT_ID,
  },
  webpack(config) {
		config.resolve.modules.push(__dirname); // 추가
		return config;
	}
}
