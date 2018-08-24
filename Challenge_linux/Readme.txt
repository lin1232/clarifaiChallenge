***********************************************************************************
Author: Yao Lin
Project name: Clarifai Challenge
Project function: 
a. Tag each image that found in an URL, and store the results in an in-memory data 
structure. 
b. Provide a html page that repeatedly reads in a string tag name and returns 
a sorted list of at most 10 of the most probable images.
************************************************************************************

************************************************************************************
How to use the program:
There are two version:windows and linux
For someone who use windows:
command: to/your/path/Challenge_windows/bin/main your_app_key model_name model_type
for example: 
./main 
use the data in the data.json, which is tagged by general-v1.3 model
./main d7d3f93929814b419b25e96869f7aefa 
use the model general-v1.3, which is the default model
./main d7d3f93929814b419b25e96869f7aefa travel-v1.0 concept
use the travel-v1.0 model

************************************************************************************

************************************************************************************
Function of all package:

application: Initialize an application with my API key and get a model, the default model
is general-v1.3.

main: It is the main function

search: Build a Tf-idf model and return the top ten links.

util: You don't need to read this package, it is not important. I write some tools in this 
package

webclient: The package is used to get or post data to network

webpage: The package is used to create a html file by a html template.

webpage.html: It is a html template

data.json: I store all url of images and its tags in this file, but as the requirement　is
storing the results in an in-memory data structure, so it is not used. If you want to use
it, you can use the Commented-Out Code in the main function.
**************************************************************************************

***********************************************************************************
Eplanation of inefficiencies:
a. This program only support public models, due to I don't have any other model.
b. I just test general-v1.3 because my usage is limited, I don't have enough usage 
to test all models. But I think other public models can also work.
c. The program will only return the rul of images that probability is bigger than zero. So
sometimes, the number of result is less than ten, as the probability of others images is 0.0. 
****************************************************************************************

****************************************************************************************
Possible improvements in the future:
a. I don't think storing the results of tagged image in an in-memory data structure is 
a good idea, if possible, it is better to store the results in a database.
b. This program only support public models, but I think I can add private model to this 
program in the future.
c. The search strategy is very simple, if possible, I think it is better to 
add some algorithm of natural language processing
***************************************************************************************