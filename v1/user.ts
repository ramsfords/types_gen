/* eslint-disable */

export const protobufPackage = "v1";

export interface user {
  /** @gotags: dynamodbav:"name,omitempty" */
  name: string;
  /** @gotags: dynamodbav:"userName,omitempty" */
  userName: string;
  /** @gotags: dynamodbav:"password,omitempty" */
  password: string;
  /** @gotags: dynamodbav:"confirmPassword,omitempty" */
  confirmPassword: string;
  /** @gotags: dynamodbav:"origin,omitempty" */
  origin: string;
  /** @gotags: dynamodbav:"emailVisibility,omitempty" */
  emailVisibility: boolean;
  /** @gotags: dynamodbav:"type,omitempty" */
  type: string;
  /** @gotags: dynamodbav:"id,omitempty" */
  id: string;
  /** @gotags: dynamodbav:"created,omitempty" */
  created: string;
  /** @gotags: dynamodbav:"updated,omitempty" */
  updated: string;
  /** @gotags: dynamodbav:"verified,omitempty" */
  verified: string;
  /** @gotags: dynamodbav:"avatar,omitempty" */
  avatar: string;
  /** @gotags: dynamodbav:"lastResetSentAt,omitempty" */
  lastResetSentAt: string;
  /** @gotags: dynamodbav:"lastVerificationSentAt,omitempty" */
  lastVerificationSentAt: string;
  /** @gotags: dynamodbav:"passwordHash,omitempty" */
  passwordHash: string;
  /** @gotags: dynamodbav:"tokenKey,omitempty" */
  tokenKey: string;
  /** @gotags: dynamodbav:"email,omitempty" */
  email: string;
}

export interface LoginUser {
  email: string;
  password: string;
}

export interface pocketCollection {
  name: string;
  type: string;
}
