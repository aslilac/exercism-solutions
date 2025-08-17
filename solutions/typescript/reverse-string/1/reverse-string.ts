export function reverse(text: string) {
  let txet = '';
  for (let i = text.length - 1; i >= 0; i--) {
    txet += text[i];
  }
  return txet;
}
