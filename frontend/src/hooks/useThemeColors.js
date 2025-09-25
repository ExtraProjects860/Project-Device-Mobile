import resolveConfig from "tailwindcss/resolveConfig";
import { useTheme } from "../context/ThemeContext.js";
import tailwindConfig from "../../tailwind.config.js";

const fullConfig = resolveConfig(tailwindConfig);
const colors = fullConfig.theme.colors;

export function useThemeColors() {
  const { isThemeDark } = useTheme();

  return {
    primary: isThemeDark ? colors.dark["text-primary"] : colors.light.primary,

    header: isThemeDark
      ? colors.dark["text-primary"]
      : colors.light["text-inverted"],

    switch: {
      track: {
        false: colors.dark["text-secondary"],
        true: colors.light.primary,
      },
      thumb: isThemeDark ? colors.dark.primary : colors.light.card,
    },
  };
}
