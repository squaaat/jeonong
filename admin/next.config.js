// next.config.js에서 promise, async function 안먹음.
// 거지같은 next.config.js
const env = {
  J_ENV: process.env.J_ENV,
  J_CICD: process.env.J_CICD,
  ADMIN_AWS_ACM_ARN: process.env.ADMIN_AWS_ACM_ARN,
  ADMIN_AWS_ROLE_ARN: process.env.ADMIN_AWS_ROLE_ARN,
}

console.log(env)

module.exports = {
  env,
  webpack(config) {
		config.resolve.modules.push(__dirname); // absolute paht를 위해

    if (!process.env.BUNDLE_AWS_SDK) { // aws-sdk bundling을 위해
      config.externals = config.externals || [];
      config.externals.push({ "aws-sdk": "aws-sdk" });
    } else {
      console.warn("Bundling aws-sdk. Only doing this in development mode");
    }

		return config;
	},
}
