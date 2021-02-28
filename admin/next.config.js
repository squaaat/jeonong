// next.config.js에서 promise, async function 안먹음.
// 거지같은 next.config.js
const env = {
  J_ENV: process.env.J_ENV,
  J_CICD: process.env.J_CICD,
  GOOGLE_OAUTH_CLIENT_ID: process.env.GOOGLE_OAUTH_CLIENT_ID,
  ADMIN_AWS_ACM_ARN: process.env.ADMIN_AWS_ACM_ARN,
}

console.log(env)

module.exports = {
  env,
  webpack(config) {
		config.resolve.modules.push(__dirname); // 추가
		return config;
	}
}
