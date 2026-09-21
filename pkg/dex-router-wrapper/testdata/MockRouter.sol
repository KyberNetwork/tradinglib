// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// MockRouter is a test-only fixture standing in for a real DEX router. It is
// compiled and embedded solely by DexRouterWrapper's Go tests, to exercise
// wrap() end-to-end against a real EVM (go-ethereum's simulated backend)
// without depending on any live network.
//
// To regenerate the compiled artifacts checked in alongside this file:
//   npx solc --optimize --abi --bin --bin-runtime MockRouter.sol -o build/
// then copy build/MockRouter.bin to MockRouter.bin (only the creation
// bytecode is needed here, since the mock is always deployed for real by the
// test, never injected via a state override).
contract MockRouter {
    error MockRouterFail(string reason);

    // swap sends amount of the received native ETH on to recipient,
    // simulating a router that swaps into a native-ETH output.
    function swap(address payable recipient, uint256 amount) external payable {
        require(msg.value >= amount, "MockRouter: insufficient value");
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = recipient.call{value: amount}("");
        require(success, "MockRouter: transfer failed");
    }

    // fail always reverts, so tests can assert DexRouterWrapper.wrap bubbles
    // up the inner revert reason instead of swallowing it.
    function fail() external pure {
        revert MockRouterFail("mock router failure");
    }
}
