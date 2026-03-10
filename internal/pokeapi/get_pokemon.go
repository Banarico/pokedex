package pokeapi

import (
        "encoding/json"
        "io"
        "net/http"
)

type Pokemon struct {
    Abilities []struct {
        Ability    struct {
            Name   string `json:"name"`
            URL    string `json:"url"`
        } `json:"ability"`
        Hidden     bool `json:"is_hidden"`
        Slot       int `json:"slot"`
    } `json:"abilities"`
    BaseExp   int `json:"base_experience"`
    Cries     struct {
        Latest string `json:"latest"`
        Legacy string `json:"legacy"`
    } `json:"cries"`
    Forms     []struct {
        Name  string `json:"name"`
        URL   string `json:"url"`
    } `json:"forms"`
    Indices   []struct {
        Index int `json:"game_index"`
        Version struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"version"`
    } `json:"game_indices"`
    Height    int `json:"height"`
    HeldItems []struct {
        Item    struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"item"`
        VersionDet    []struct {
            Rarity    int `json:"rarity"`
            Version    struct {
                Name  string `json:"name"`
                URL   string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"held_items"`
    ID        int `json:"id"`
    Default   bool `json:"is_default"`
    Encounters string `json:"location_area_encounters"`
    Moves     []struct {
        Move  struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"move"`
        VerGroup []struct {
            LearnedAt int `json:"level_learned_at"`
            Method    struct {
                Name  string `json:"name"`
                URL   string `json:"url"`
            } `json:"move_learn_method"`
            Order     int `json:"order"`
            Group     struct {
                Name  string `json:"name"`
                URL   string `json:"url"`
            } `json:"version_group"`
        } `json:"version_group_details"`
    } `json:"moves"`
    Name      string `json:"name"`
    Order     int `json:"order"`
    PastAbl   []struct {
        Abilities    []struct {
            Ability  struct {
                Name  string `json:"name"`
                URL   string `json:"url"`
            } `json:"ability"`
            Hidden   bool `json:"is_hidden"`
            Slot     int `json:"slot"`
        } `json:"abilities"`
        Generation   struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"generation"`
    } `json:"past_abilities"`
    PastSts   []struct {
        Generation    struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"generation"`
        Stats     []struct {
            BaseStat    int `json:"base_stat"`
            Effort      int `json:"effort"`
            Stat        struct {
                Name  string `json:"name"`
                URL   string `json:"url"`
            } `json:"stat"`
        } `json:"stats"`
    } `json:"past_stats"`
    PastTyp   []struct {
    } `json:"past_types"`
    Species   struct {
        Name  string `json:"name"`
        URL   string `json:"url"`
    } `json:"species"`
    Sprites   struct {
        BackDef    string `json:"back_default"`
        BackFem    string `json:"back_female"`
        BackShi    string `json:"back_shiny"`
        BackShiFem string `json:"back_shiny_female"`
        FrontDef   string `json:"front_default"`
        FrontFem   string `json:"front_female"`
        FrontShi   string `json:"front_shiny"`
        FrontShiFem string `json:"front_shiny_female"`
        Other      struct {
            DreamWorld    struct {
                FrontDef   string `json:"front_default"`
                FrontFem   string `json:"front_female"`
            } `json:"dream_world"`
            Home          struct {
                FrontDef   string `json:"front_default"`
                FrontFem   string `json:"front_female"`
                FrontShi   string `json:"front_shiny"`
                FrontShiFem string `json:"front_shiny_female"`
            } `json:"home"`
            OfficialArt   struct {
                FrontDef   string `json:"front_default"`
                FrontShi   string `json:"front_shiny"`
            } `json:"official-artwork"`
            Showdown      struct {
                BackDef    string `json:"back_default"`
                BackFem    string `json:"back_female"`
                BackShi    string `json:"back_shiny"`
                BackShiFem string `json:"back_shiny_female"`
                FrontDef   string `json:"front_default"`
                FrontFem   string `json:"front_female"`
                FrontShi   string `json:"front_shiny"`
                FrontShiFem string `json:"front_shiny_female"`
            } `json:"showdown"`
        } `json:"other"`
        Versions   struct {
            GenI   struct {
                RedBlue    struct {
                    BackDef    string `json:"back_default"`
                    BackGray   string `json:"back_gray"`
                    BackTran   string `json:"back_transparent"`
                    FrontDef   string `json:"front_default"`
                    FrontGray  string `json:"front_gray"`
                    FrontTran  string `json:"front_transparent"`
                } `json:"red-blue"`
                Yellow    struct {
                    BackDef    string `json:"back_default"`
                    BackGray   string `json:"back_gray"`
                    BackTran   string `json:"back_transparent"`
                    FrontDef   string `json:"front_default"`
                    FrontGray  string `json:"front_gray"`
                    FrontTran  string `json:"front_transparent"`
                } `json:"yellow"`
            } `json:"generation-i"`
            GenII  struct {
                Crystal    struct {
                    BackDef    string `json:"back_default"`
                    BackShi    string `json:"back_shiny"`
                    BackShiTran string `json:"back_shiny_transparent"`
                    BackTran   string `json:"back_transparent"`
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiTran  string `json:"front_shiny_transparent"`
                    FrontTran  string `json:"front_transparent"`
                } `json:"crystal"`
                Gold       struct {
                    BackDef    string `json:"back_default"`
                    BackShi    string `json:"back_shiny"`
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                    FrontTran  string `json:"front_transparent"`
                } `json:"gold"`
                Silver     struct {
                    BackDef    string `json:"back_default"`
                    BackShi    string `json:"back_shiny"`
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                    FrontTran  string `json:"front_transparent"`
                } `json:"silver"`
            } `json:"generation-ii"`
            GenIII struct {
                Emerald    struct {
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                } `json:"emerald"`
                FireLeaf   struct {
                    BackDef    string `json:"back_default"`
                    BackShi    string `json:"back_shiny"`
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                } `json:"firered-leafgreen"`
                RubySaph   struct {
                    BackDef    string `json:"back_default"`
                    BackShi    string `json:"back_shiny"`
                    FrontDef   string `json:"front_default"`
                    FrontShi   string `json:"front_shiny"`
                } `json:"ruby-sapphire"`
            } `json:"generation-iii"`
            GenIV  struct {
                DiamPearl    struct {
                    BackDef    string `json:"back_default"`
                    BackFem    string `json:"back_female"`
                    BackShi    string `json:"back_shiny"`
                    BackShiFem string `json:"back_shiny_female"`
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"diamond-pearl"`
                HeartSoul    struct {
                    BackDef    string `json:"back_default"`
                    BackFem    string `json:"back_female"`
                    BackShi    string `json:"back_shiny"`
                    BackShiFem string `json:"back_shiny_female"`
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"heartgold-soulsilver"`
                Platinum     struct {
                    BackDef    string `json:"back_default"`
                    BackFem    string `json:"back_female"`
                    BackShi    string `json:"back_shiny"`
                    BackShiFem string `json:"back_shiny_female"`
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"platinum"`
            } `json:"generation-iv"`
            GenIX  struct {
                ScarViol    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                } `json:"scarlet-violet"`
            } `json:"generation-ix"`
            GenV   struct {
                BlackWhite    struct {
                    Animated    struct {
                        BackDef    string `json:"back_default"`
                        BackFem    string `json:"back_female"`
                        BackShi    string `json:"back_shiny"`
                        BackShiFem string `json:"back_shiny_female"`
                        FrontDef   string `json:"front_default"`
                        FrontFem   string `json:"front_female"`
                        FrontShi   string `json:"front_shiny"`
                        FrontShiFem string `json:"front_shiny_female"`
                    } `json:"animated"`
                    BackDef    string `json:"back_default"`
                    BackFem    string `json:"back_female"`
                    BackShi    string `json:"back_shiny"`
                    BackShiFem string `json:"back_shiny_female"`
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"black-white"`
            } `json:"generation-v"`
            GenVI  struct {
                OmegaAlpha    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"omegaruby-alphasapphire"`
                XY            struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"x-y"`
            } `json:"generation-vi"`
            GenVII struct {
                Icons    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                } `json:"icons"`
                UltraSunMoon    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                    FrontShi   string `json:"front_shiny"`
                    FrontShiFem string `json:"front_shiny_female"`
                } `json:"ultra-sun-ultra-moon"`
            } `json:"generation-vii"`
            GenVIII struct {
                BrillShin    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                } `json:"brilliant-diamond-shining-pearl"`
                Icons    struct {
                    FrontDef   string `json:"front_default"`
                    FrontFem   string `json:"front_female"`
                } `json:"icons"`
            } `json:"generation-viii"`
        } `json:"versions"`
    } `json:"sprites"`
    Stats     []struct {
        BaseStat    int `json:"base_stat"`
        Effort      int `json:"effort"`
        Stat        struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"stat"`
    } `json:"stats"`
    Types     []struct {
        Slot    int `json:"slot"`
        Type    struct {
            Name  string `json:"name"`
            URL   string `json:"url"`
        } `json:"type"`
    } `json:"types"`
    Weight    int `json:"weight"`
}

func (c *Client) GetPokemon(arg string) (Pokemon, error) {
    url := "https://pokeapi.co/api/v2/pokemon/" + arg
    req, err := http.Get(url)
    if err != nil {
        return Pokemon{}, err
    }
    data, err := io.ReadAll(req.Body)
    if err != nil {
        return Pokemon{}, err
    }
    pkm := Pokemon{}
    err = json.Unmarshal(data, &pkm)
    if err != nil {
        return Pokemon{}, err
    }
    return pkm, nil
}
