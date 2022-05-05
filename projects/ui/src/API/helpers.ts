import { ObjectMeta } from 'proto/github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/common_pb';
import {
  ObjectRef,
  ClusterObjectRef,
} from 'proto/github.com/solo-io/skv2/api/core/v1/core_pb';
import useSWR, { Key } from 'swr';

export const host = `${
  process.env.NODE_ENV === 'production'
    ? window.location.origin
    : 'http://localhost:8090'
}`;

export function getObjectRefClassFromRefObj(
  requestMeta: ObjectRef.AsObject
): ObjectRef {
  let ref = new ObjectRef();
  ref.setName(requestMeta.name);
  ref.setNamespace(requestMeta.namespace);
  return ref;
}

export function getClusterRefClassFromClusterRefObj(
  requestMeta: ClusterObjectRef.AsObject
): ClusterObjectRef {
  let ref = new ClusterObjectRef();
  ref.setName(requestMeta.name);
  ref.setNamespace(requestMeta.namespace);
  ref.setClusterName(requestMeta.clusterName);
  return ref;
}

export function setObjectMeta(
  setFunc: (rr: ObjectMeta) => void,
  resRefObj?: ObjectMeta.AsObject
): void {
  if (resRefObj !== undefined) {
    let resRef = new ObjectMeta();
    resRef.setName(resRefObj.name);
    resRef.setNamespace(resRefObj.namespace);
    setFunc(resRef);
  }
}

export function objectMetasAreEqual(
  resRef1?:
    | ObjectRef.AsObject
    | ObjectMeta.AsObject
    | ClusterObjectRef.AsObject,
  resRef2?:
    | ObjectRef.AsObject
    | ObjectMeta.AsObject
    | ClusterObjectRef.AsObject,
  onlyCompareRefs?: boolean
) {
  if (!resRef1 || !resRef2) {
    return false;
  }

  if (!onlyCompareRefs) {
    if (
      // @ts-ignore
      (resRef1.clusterName !== undefined || // @ts-ignore
        resRef2.clusterName !== undefined) &&
      // @ts-ignore
      resRef1.clusterName !== resRef2.clusterName
    ) {
      return false;
    }
  }

  return (
    resRef1.name === resRef2.name && resRef1.namespace === resRef2.namespace
  );
}

const NORMAL_REFRESH_INTERVAL = 10000;

type useRequestOptions =
  | {
      key: Key;
      skip?: boolean;
    }
  | {
      methodDescriptorName: string;
      skip?: boolean;
    };

/**
 * Calls and returns `useSWR(...)` for the given function + arguments.
 * Generates a cached key by default from the function call.
 * @param fn The function to call.
 * @param fnArgs The function arguments for `fn`.
 * @param options
 * ```
 * {
 *  key?: 'custom-key' // Overrides the generated key.
 *  skip?: true // Sets key=null to skip the API request.
 * }
 * ```
 * @param swrConfig This is the same config that is passed into `useSWR`.
 * @returns `useSWR(...)`
 */
export function useRequest<T extends (...args: any) => Promise<any>>(
  fn: T,
  fnArgs: Parameters<T>,
  options?: useRequestOptions,
  swrConfig?: Parameters<typeof useSWR>[2]
) {
  // Set the method descriptor name to the function name if not supplied
  // TODO: If methodDescriptors are added to the API methods,
  // then options should be required and this line can be removed.
  if (options === undefined) options = { methodDescriptorName: fn.name };
  //
  // Set the key and return useSWR(...).
  let key: Key;
  if (options?.skip === true) key = null;
  else if ('key' in options) key = options.key;
  else {
    // Generates the key from the function + arguments.
    // Removes undefined/optional arguments from the end of the key.
    key = [options.methodDescriptorName, ...fnArgs] as any[];
    while (key.length > 1 && key[key.length - 1] === undefined) key.pop();
  }
  return useSWR<Awaited<ReturnType<T>>>(key, () => fn(...(fnArgs as any[])), {
    refreshInterval: NORMAL_REFRESH_INTERVAL,
    ...(swrConfig as any),
  });
}
