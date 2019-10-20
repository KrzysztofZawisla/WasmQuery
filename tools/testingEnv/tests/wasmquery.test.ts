const fs = require("fs");
const path = require("path");

let jsonWithOutput;

beforeAll((done) => {
    fs.readFile(path.join(__dirname, "../temp/outputTest.json"), {
        encoding: "utf8"
    }, (err, data) => {
        jsonWithOutput = JSON.parse(data);
        done();
    });
});

test("Web page can be loaded", () => {
    expect(jsonWithOutput.pageHasBeenLoaded).toBe(true);
});