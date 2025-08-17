export function isPangram(text: string) {
  const alphabet = new Set([
    'a',
    'b',
    'c',
    'd',
    'e',
    'f',
    'g',
    'h',
    'i',
    'j',
    'k',
    'l',
    'm',
    'n',
    'o',
    'p',
    'q',
    'r',
    's',
    't',
    'u',
    'v',
    'w',
    'x',
    'y',
    'z',
  ]);

  for (let i = 0; i < text.length; i++) {
    alphabet.delete(text[i].toLowerCase());
  }

  return alphabet.size === 0;
}
