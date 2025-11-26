/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will demonstrate if nested statements 
 */

// Variables
let userName: string = "";
let student: string = "";
let userAge: string = "";
let userAgeNumber: number = 0;

// INPUT
userName = prompt("Hello! Please state your name") || "";
student = prompt("Okay, now are you a student?") || "";

if (student == "yes" || student == "Yes") {
userAge = prompt("Okay, now please state your age") || "0";
userAgeNumber = parseInt(userAge);

  if (userAgeNumber >= 13 && userAgeNumber <= 19) {
    console.log("Student " + userName + " is a teenager")
  } else if (userAgeNumber <= 12 && userAgeNumber >= 5) {
    console.log("Student " + userName + " is a child")

  }

  }
  


if (student == "no" || student == "No") {
  userAge = prompt("Okay, now please state your age") || "";
  userAgeNumber = parseInt(userAge);

  if (userAgeNumber >= 20 && userAgeNumber <= 30) {
    console.log (userName + " is a young adult")
  } else { console.log(userName + " is in a different life stage")
    
  }
  
}

console.log("\nDone.")

