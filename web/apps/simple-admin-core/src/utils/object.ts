export function get(obj: object, path: string, defaultValue?: any) {
  const travel = (regexp: RegExp) =>
    String.prototype.split
      .call(path, regexp)
      .filter(Boolean)
      // eslint-disable-next-line unicorn/no-array-reduce
      .reduce(
        (res, key) =>
          res !== null && res !== undefined ? (res as any)[key] : res,
        obj,
      );
  const result = travel(/[,[\]]+?/) || travel(/[,[\].]+?/);
  return result === undefined || result === obj ? defaultValue : result;
}
