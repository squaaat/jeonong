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

// module.exports = async function() {
//   const target = {
//     env: {}
//   }
//   var ssm = new AWS.SSM();

//   var options = {
//     Name: '/jeonong/jeonong-api/alpha/application.yml', /* required */
//     WithDecryption: true
//   };

//   const data = await ssm.getParameter(options).promise();
//   target.env.HELLO = data.Parameter.Value

//   return target
// }

// module.exports = new Promise((resolve, reject) => {
//   const target = {
//     env: {}
//   }
//   var ssm = new AWS.SSM();

//   var options = {
//     Name: '/jeonong/jeonong-api/alpha/application.yml', /* required */
//     WithDecryption: true
//   };

//   const parameterPromise = ssm.getParameter(options).promise();
//   parameterPromise.then(function(data, err) {
//     if (err) reject(err)
//     target.env.HELLO = data.Parameter.Value
//     resolve(target)
//   });
// })

// module.exports = function(...args) {
//   let original = require('./next.config.original.1612888225653.js');
//   const finalConfig = {};
//   const target = { target: 'serverless' };
//   if (typeof original === 'function' && original.constructor.name === 'AsyncFunction') {
//     // AsyncFunctions will become promises
//     original = original(...args);
//   }
//   if (original instanceof Promise) {
//     // Special case for promises, as it's currently not supported
//     // and will just error later on
//     return original
//       .then((originalConfig) => Object.assign(finalConfig, originalConfig))
//       .then((config) => Object.assign(config, target));
//   } else if (typeof original === 'function') {
//     Object.assign(finalConfig, original(...args));
//   } else if (typeof original === 'object') {
//     Object.assign(finalConfig, original);
//   }
//   Object.assign(finalConfig, target);
//   return finalConfig;
// }