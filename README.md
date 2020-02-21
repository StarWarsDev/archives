# Star Wars Legion Archives
A read-only GraphQL API for querying data about Star Wars Legion.

## Access
The API is read only and does **not** require any authentication.

You can find the GraphQL Playground here: https://sw-legion-archives.herokuapp.com and the graphql
endpoint can be found at https://sw-legion-archives.herokuapp.com/graphql.

## Data
The data backing this API comes directly from the `src/data.json` file from
[@NicholasCBrown/legion-HQ-2.0](https://github.com/NicholasCBrown/legion-HQ-2.0).

## Using the `query` argument
Unless noted below, all `query` arguments allow for searching with regular expressions.

### Query syntax
All `query` argument values follow the same pattern where the field is placed to the left and the value is placed on the
right and are separated using a colon.

**Example:** Query keywords where the description contains the word `rally`.  By prefixing the search term with `(?i)`
the search becomes case insensitive.
```graphql
{
  keywords(query: "description:(?i)rally") {
    name
    description
  }
}
```

### Fields that support *only* exact match values
The following fields will only match on exact values and will ignore regular expressions.

- commands
    - id
    - pips
- units
    - id
    - unique
- upgrades
    - id
    - unique