const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
	let connections;
	try {
		connections = await ddb
			.scan({ TableName: process.env.RTO_OFFICE_TABLE_ARN })
			.promise();
	} catch (err) {
		return {
			statusCode: 500,
			message: "Internal Server Error : Error in fetching from DB.",
		};
	}
	const callbackAPI = new AWS.ApiGatewayManagementApi({
		apiVersion: "2018-11-29",
		endpoint:
			event.requestContext.domainName + "/" + event.requestContext.stage,
	});

	const message = JSON.parse(event.body).message;
	console.log(JSON.parse(event.body));

	const sendMessages = connections.Items.map(async ({ connectionId }) => {
		if (connectionId !== event.requestContext.connectionId) {
			try {
				await callbackAPI
					.postToConnection({ ConnectionId: connectionId, Data: message })
					.promise();
			} catch (e) {
				console.log(e);
			}
		}
	});

	try {
		await Promise.all(sendMessages);
	} catch (e) {
		console.log(e);
		return {
			statusCode: 500,
		};
	}

	return { statusCode: 200 };
};
