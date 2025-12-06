const fs = require('fs').promises;
const path = require('path');
const outputSwift = path.join(__dirname, 'outputSwift.txt');
const outputCorrect = path.join(__dirname, 'outputCorrect.txt');

(async () => {
    const swiftData = await fs.readFile(outputSwift, {encoding: 'utf-8'});
    const swiftJson = JSON.parse(swiftData);

    const correctData = await fs.readFile(outputCorrect, {encoding: 'utf-8'});
    const correctJson = JSON.parse(correctData);
    
    swiftJson.forEach((swift, index) => {
        const correct = correctJson[index];
        if (swift !== correct) {
            console.log(`Index: ${index} - Swift: ${swift} - Correct: ${correct}`);
        }
    });
})();
