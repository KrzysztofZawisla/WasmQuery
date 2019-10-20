import * as express from "express";
import * as path from "path";
import * as fs from "fs";

const app = express();
const port: number = 4000;

app.use(express.static(path.join(__dirname, "public")));

app.get("/", (req, res) => {
	if(fs.existsSync(path.join(__dirname, "public/index.html"))) {
		res.sendFile(path.join(__dirname, "public/index.html"));
		res.header(200);
	} else {
		res.header(404)
	}
});

app.listen(port, () => {
	console.log("Server running on port: "+port);
});
