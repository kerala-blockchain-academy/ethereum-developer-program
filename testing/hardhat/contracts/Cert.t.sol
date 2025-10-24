// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Cert} from "./Cert.sol";
import {Test} from "forge-std/Test.sol";

contract CertTest is Test {
    Cert cert;

    function setUp() public {
        cert = new Cert();
    }

    function test_Issue() public {
        vm.expectEmit(true, false, false, false);
        emit Cert.Issued("MBCC", 21, "2025");
        cert.issue(21, "Shin", "MBCC", "S", "2025");
    }

    function test_Certificates() public {
        cert.issue(419, "McQueen", "MBCC", "A", "2023");
        string memory name;
        string memory course;
        string memory grade;
        string memory date;
        (name, course, grade, date) = cert.Certificates(419);
        Cert.Certificate memory cc = Cert.Certificate(
            "McQueen",
            "MBCC",
            "A",
            "2023"
        );
        assertEq(name, cc.name);
        assertEq(course, cc.course);
        assertEq(grade, cc.grade);
        assertEq(date, cc.date);
    }

    function test_Revert() public {
        vm.expectRevert("Access Denied");
        vm.prank(address(0));
        cert.issue(270, "Che", "MBCC", "B", "2024");
    }
}
