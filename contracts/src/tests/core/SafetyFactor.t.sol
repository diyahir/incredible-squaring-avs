//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import "../../static/Structs.sol";
import "../../AVSReservesManager.sol";
import "../../SafetyFactorOracle.sol";

// Mocks

contract SafetyFactorOracleTests is Test {
    AVSReservesManager public reservesManager;
    SafetyFactorOracle public safetyFactorOracle;
    SafetyFactorConfig public safetyFactorConfig;

    address public anzenTaskManager;
    address public anzenGov;

    address public avsGov;
    address public avsServiceManager;
    uint32 public avsId;

    address[] public rewardTokens;
    uint256[] public initialTokenFlows;

    function setUp() public {
        anzenTaskManager = address(0x123);
        anzenGov = address(0x456);

        avsId = uint32(1);
        avsGov = address(0x789);
        avsServiceManager = address(0x222);

        safetyFactorOracle = new SafetyFactorOracle(anzenTaskManager, anzenGov);
    }

    function test_addProtocol() public virtual {
        vm.prank(anzenGov);
        safetyFactorOracle.addProtocol(avsId, avsGov);
    }

    function test_reverts_addProtocol_notAnzenGovernance() public virtual {
        vm.expectRevert();
        safetyFactorOracle.addProtocol(avsId, avsGov);
    }

    function test_removeProtocol() public virtual {
        vm.prank(anzenGov);
        safetyFactorOracle.addProtocol(avsId, avsGov);
        vm.prank(anzenGov);
        safetyFactorOracle.removeProtocol(avsId);
    }

    function test_reverts_removeProtocol_notAnzenGovernance() public virtual {
        vm.expectRevert();
        safetyFactorOracle.removeProtocol(avsId);
    }

    function test_setSafetyFactor(int256 _sf) public virtual {
        avsId = uint32(1);
        vm.prank(anzenGov);
        safetyFactorOracle.addProtocol(avsId, avsGov);

        vm.prank(anzenTaskManager);
        safetyFactorOracle.setSafetyFactor(avsId, _sf);

        SafetyFactorSnapshot memory snapshot = safetyFactorOracle.getSafetyFactor(avsId);
        assertEq(snapshot.safetyFactor, _sf);
        assertEq(snapshot.timestamp, block.timestamp);
    }

    function test_reverts_setSafetyFactor_notAnzenTaskManager(int256 _sf) public virtual {
        vm.expectRevert();
        safetyFactorOracle.setSafetyFactor(avsId, _sf);
    }
}
