# ////////////////////////////////////////////////////////////////////////////
# //
# //  D SOFTWARE INCORPORATED
# //  Copyright 2007-2014 D Software Incorporated
# //  All Rights Reserved.
# //
# //  NOTICE: D Software permits you to use, modify, and distribute this file
# //  in accordance with the terms of the license agreement accompanying it.
# //
# //  Unless required by applicable law or agreed to in writing, software
# //  distributed under the License is distributed on an "AS IS" BASIS,
# //  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# //
# ////////////////////////////////////////////////////////////////////////////
#
# This is a sample of what can be done by using API with Zephyr for JIRA through the Python coding language.
# The goal of the sample is to create a new test cycle and add existing test issues into it
#
# IDLE IDE - Version: 2.7.5
# Python - Version: 2.7.5
#
# Author: Daniel Gannon, Technical Support Analyst, D Software Inc.

# Import from libraries 'urllib2', 'json', and 'base64'
from urllib2 import Request, urlopen
from json import dumps, load
from base64 import b64encode

# Variables used:
# 'username' and 'password' variables for encoding into header below
username = 'test.manager'
password = 'test.manager'
# New Test Issue Variables:
issueProjectKey = "WRAEC"
issueSummary = "Login Service Test"
issueDescription = "Login Credential Combination Test"
issueTypeName = "Test"
# New Test Step Variables:
testStep = "Validate Username/Password Combination"
stepData = "Username/Password"
stepResult = "Successful validation message"

# 'baseURL' holds basic data for JIRA server that will be used in API calls
# 'createTestIssueURL' hold URL sent to JIRA to create a new Issue
# 'createTestStepURL' hold URL sent to ZAPI to add a test step to a Test issue
baseURL = 'http://localhost:8080'
createTestIssueURL = baseURL + "/rest/api/2/issue/"
createTestStepURL = baseURL + "/rest/zapi/latest/teststep/"

# 'headers' holds information to be sent with the JSON data set
# Initialized with Auth and Content-Type data
# Authorization header uses base64 encoded username and password string
headers = {"Authorization": " Basic " + b64encode(username + ":" + password), "Content-Type": "application/json"}

# ///// Create Test Issue /////
# Create a new Zephyr for JIRA Test issue in specified project

print "Creating Test Issue..."

newTestValues = dumps({
    "fields": {
        "project": {
            "key": issueProjectKey
        },
        "summary": issueSummary,
        "description": issueDescription,
        "issuetype": {
            "name": issueTypeName
        }
    }
})

request = Request(createTestIssueURL, data=newTestValues, headers=headers)
js_res = urlopen(request)
objResponse = load(js_res)
newTestId = objResponse['id']

print "Test Issue Created! New ID is: " + newTestId

# ///// Create Test Step in new Test Issue /////
# Use new Test issue ID created to add a new test step
# 'newTestId' is the ID of the Test issue created above and returned from JIRA

print "Adding New Step To Newly Created Test Issue"

newTestStepValues = dumps ({
    "step": testStep,
    "data": stepData,
    "result": stepResult
})

createTestStepURL = createTestStepURL + newTestId

request = Request(createTestStepURL, data=newTestStepValues, headers=headers)
urlopen(request)

print "Step Addition Completed!"
