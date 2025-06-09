// *******************************************************************************
//  Contributors: BMC Helix, Inc.
//
//  (c) Copyright 2020-2025 BMC Helix, Inc.
//
//* SPDX-License-Identifier: Apache-2.0
// *******************************************************************************

const userIdHeader = process.env.UID_HTTP_HEADER || 'X-Credential-Identifier';

function userHeader(r) {
    // Extract the username from the request header
    return r.headersIn[userIdHeader];
}

export default {userHeader}