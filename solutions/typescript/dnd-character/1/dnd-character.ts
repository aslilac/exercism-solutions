export class DnDCharacter {
  strength = DnDCharacter.generateAbilityScore();
  dexterity = DnDCharacter.generateAbilityScore();
  constitution = DnDCharacter.generateAbilityScore();
  intelligence = DnDCharacter.generateAbilityScore();
  wisdom = DnDCharacter.generateAbilityScore();
  charisma = DnDCharacter.generateAbilityScore();
  hitpoints = 10 + DnDCharacter.getModifierFor(this.constitution);
  
  public static generateAbilityScore(): number {
    let lowest = 6;
    let total = 0;

    for (let i = 0; i < 3; i++) {
      let next = 1 + (Math.random() * 6) << 0;
      lowest = Math.min(lowest, next);
      total += next;
    }

    return total - lowest;
  }

  public static getModifierFor(abilityValue: number): number {
    return Math.floor((abilityValue - 10) / 2);
  }
}
