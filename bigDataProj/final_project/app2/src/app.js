'use strict';
const http = require('http');
var assert = require('assert');
const express= require('express');
const app = express();
const mustache = require('mustache');
const filesystem = require('fs');
const url = require('url');
const port = Number(process.argv[2]);

const hbase = require('hbase')
//host '127.0.0.1' port = 3666
var hclient = hbase({ host: process.argv[3], port: Number(process.argv[4])})

function rowToMap(row) {
	var stats = {}
	row.forEach(function (item) {
		stats[item['column']] = Number(item['$'])
	});
	return stats;
}

hclient.table('CCmortgage_v3').row('60004').get((error, value) => {
	console.info(rowToMap(value))
	console.info(value)
})


//tell app.js to go check in mortgage.html in the public directory
app.use(express.static('public'));
app.get('/mortgage_results.html',function (req, res) {
    const zipcode=req.query['zipcode'];
    console.log(zipcode);
	//instead of .get .scan
	hclient.table('CCmortgage_v3').row(zipcode).get(function (err, cells) {
		const mortgageDeathInfo = rowToMap(cells);
		console.log(mortgageDeathInfo)
		function average_mortgage_deaths(zipcode) {
			var num_homes = mortgageDeathInfo["Cook County:" + zipcode + "homes"];
			var sum_mortgages = mortgageDeathInfo["Cook County:" + zipcode + "spent on mortages"];
			var sum_accident = mortgageDeathInfo[zipcode];
			var sum_homicide = mortgageDeathInfo[zipcode];
			var sum_natural = mortgageDeathInfo[zipcode];
			var sum_suicide = mortgageDeathInfo[zipcode];
			var sum_undetermined = mortgageDeathInfo[zipcode];
			var sum_deaths = mortgageDeathInfo[zipcode];
			if(num_homes == 0 || sum_deaths == 0)
				return " - ";
			return (sum_mortgages/num_homes).toFixed(1), (sum_accident/sum_deaths).toFixed(1), (sum_homicide/sum_deaths).toFixed(1),
				(sum_natural/sum_deaths).toFixed(1), (sum_suicide/sum_deaths).toFixed(1), (sum_undetermined/sum_deaths).toFixed(1); /* One decimal place */
		}

		var template = filesystem.readFileSync("result.mustache").toString();
		var html = mustache.render(template,  {
			zipcode : req.query['Zip Code'],
			averageMortgage : average_mortgage_deaths("Average Mortgage"),
			perDeathAccident : average_mortgage_deaths("% Death Accident"),
			perDeathHomicide : average_mortgage_deaths("% Death Homicide"),
			perDeathNatural : average_mortgage_deaths("% Death Natural"),
			perDeathSuicide : average_mortgage_deaths("% Death Suicide"),
			perDeathUndetermined : average_mortgage_deaths("% Death Undetermined"),
		});
		res.send(html);
	});
});

/* Send simulated new mortgages to kafka */
 var kafka = require('kafka-node');
 var Producer = kafka.Producer;
 var KeyedMessage = kafka.KeyedMessage;
 var kafkaClient = new kafka.KafkaClient({kafkaHost: process.argv[5]});
 var kafkaProducer = new Producer(kafkaClient);


 app.get('/new_mortgage.html',function (req, res) {
 	var zipcode = req.query['5 digit zip code'];
 	var sum_mortgages = req.query['total sum of all mortgages'];
 	var num_homes = req.query['total number of mortgages'];
 	var sum_accident = req.query['total number of accidental deaths'];
 	var sum_homicide = req.query['total number of homicides'];
 	var sum_natural = req.query['total number of natural deaths'];
 	var sum_suicide = req.query['total number of suicides'];
 	var sum_undetermined = req.query['total number of undetermined deaths'];
 	var sum_deaths = req.query['total number of deaths, any cause'];
 	var report = {
 		zipcode : zipcode,
 		sum_mortgages : sum_mortgages,
 		num_homes : num_homes,
 		sum_accident : sum_accident,
 		sum_homicide : sum_homicide,
 		sum_natural : sum_natural,
 		sum_suicide : sum_suicide,
 		sum_undetermined : sum_undetermined,
 		sum_deaths : sum_deaths,
 	};

 	kafkaProducer.send([{ topic: 'mortgage-reports', messages: JSON.stringify(report)}],
 		function (err, data) {
 			console.log(report);
 			res.redirect('submit_mortgage.html');
 		});
 });

app.listen(port);
