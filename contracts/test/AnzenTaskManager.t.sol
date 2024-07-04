// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/AnzenServiceManager.sol" as anzensm;
import {AnzenTaskManager} from "../src/AnzenTaskManager.sol";
import {SafetyFactorOracle} from "../src/SafetyFactorOracle.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract AnzenTaskManagerTest is BLSMockAVSDeployer {
    anzensm.AnzenServiceManager sm;
    anzensm.AnzenServiceManager smImplementation;
    AnzenTaskManager tm;
    AnzenTaskManager tmImplementation;
    SafetyFactorOracle sfo;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator = address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator = address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        // First, deploy the SafetyFactorOracle contract.
        sfo = new SafetyFactorOracle(address(0), address(0));

        tmImplementation =
            new AnzenTaskManager(anzensm.IRegistryCoordinator(address(registryCoordinator)), TASK_RESPONSE_WINDOW_BLOCK);

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = AnzenTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator,
                        generator,
                        address(sfo)
                    )
                )
            )
        );
    }

    function testCreateNewOraclePullTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        tm.createNewOraclePullTask(0, 2, 0, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
