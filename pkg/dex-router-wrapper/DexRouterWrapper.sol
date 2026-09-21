// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// DexRouterWrapper forwards an arbitrary call to a DEX router and reports the
// resulting output-token amount and gas used. It is meant to be run via
// eth_call/eth_estimateGas with a state override (see the Go package's
// DeployedBytecode/StateOverride helpers) rather than deployed on-chain, so a
// single caller can quote any router's swap without holding an allowance or
// funding the wrapper itself.
//
// To regenerate the compiled artifacts checked in alongside this file:
//   npx solc --optimize --abi --bin --bin-runtime DexRouterWrapper.sol -o build/
// then copy build/DexRouterWrapper.abi to DexRouterWrapper.abi.json,
// build/DexRouterWrapper.bin to DexRouterWrapper.bin, and
// build/DexRouterWrapper.bin-runtime to DexRouterWrapper.bin-runtime.
contract DexRouterWrapper {
    // NATIVE is the sentinel address used to mean "the native chain asset"
    // wherever an ERC20 token address is otherwise expected, matching the
    // convention used by 1inch, Paraswap, and most other DEX aggregators.
    address constant NATIVE = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    // wrap calls target with data, forwarding any msg.value (needed when the
    // router call itself takes native ETH as input), and returns the change
    // in outputToken balance of recipient plus the gas the call consumed.
    //
    // recipient is taken explicitly rather than assumed to be address(this)
    // or msg.sender: the router calldata built off-chain may already encode
    // an arbitrary recipient independent of who calls target.
    function wrap(
        address target,
        bytes calldata data,
        address outputToken,
        address recipient
    ) external payable returns (uint256 returnAmount, uint256 gasUsed) {
        uint256 gasStart = gasleft();
        uint256 balanceBefore = _balanceOf(outputToken, recipient);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) = target.call{value: msg.value}(data);
        if (!success) {
            _bubbleRevert(result);
        }

        uint256 balanceAfter = _balanceOf(outputToken, recipient);
        returnAmount = balanceAfter - balanceBefore;
        gasUsed = gasStart - gasleft();
    }

    function _balanceOf(address token, address account) private view returns (uint256) {
        if (token == NATIVE) {
            return account.balance;
        }
        (bool success, bytes memory result) = token.staticcall(
            abi.encodeWithSignature("balanceOf(address)", account)
        );
        require(success && result.length >= 32, "DexRouterWrapper: balanceOf failed");
        return abi.decode(result, (uint256));
    }

    function _bubbleRevert(bytes memory result) private pure {
        if (result.length > 0) {
            // solhint-disable-next-line no-inline-assembly
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
        revert("DexRouterWrapper: target call reverted");
    }
}
