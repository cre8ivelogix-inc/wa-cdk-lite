#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { FunctionalTestsStack } from '../lib/functional-tests-stack';

const app = new cdk.App();

console.log("Executing Functional Tests...");

new FunctionalTestsStack(app, 'FunctionalTestsStack');