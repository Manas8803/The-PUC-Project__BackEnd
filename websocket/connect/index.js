const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
	const rtoOfficeName = event.queryStringParameters
		? event.queryStringParameters.office_name
		: null;
	console.log(rtoOfficeName);
	if (!rtoOfficeName) {
		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Connected successfully" }),
		};
	}

	try {
		console.log("In Connect: " + rtoOfficeName);
		await ddb
			.put({
				TableName: process.env.RTO_OFFICE_TABLE_ARN,
				Item: {
					office_name: rtoOfficeName,
					connectionId: event.requestContext.connectionId,
				},
			})
			.promise();
		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Connected successfully" }),
		};
	} catch (err) {
		console.error("Error connecting:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error connecting" }),
		};
	}
};
