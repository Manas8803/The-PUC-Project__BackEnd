const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {

	const rtoOfficeName = event["queryStringParameters"]
		? event["queryStringParameters"].office_name
		: null;

	try {
		console.log("In Disconnect: " + rtoOfficeName);
		await ddb
			.delete({
				TableName: process.env.USER_TABLE_ARN,
				Key: {
					office_name: rtoOfficeName,
				},
			})
			.promise();

		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Disconnected successfully" }),
		};
	} catch (err) {
		console.error("Error disconnecting:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error disconnecting" }),
		};
	}
};
