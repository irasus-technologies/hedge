// *******************************************************************************
//  Contributors: BMC Helix, Inc.
//
//  (c) Copyright 2020-2025 BMC Helix, Inc.
//
//* SPDX-License-Identifier: Apache-2.0
// *******************************************************************************

function getAuth(r) {
    let external_auth = process.env["IS_EXTERNAL_AUTH"];
    return external_auth !== undefined ? external_auth : "false";
}

export default {getAuth}