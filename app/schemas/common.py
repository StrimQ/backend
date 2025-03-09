from fastapi import Query


def EntityNameQuery(entity: str):
    return Query(
        min_length=3,
        max_length=50,
        pattern=r"[\w\- ]+",
        description=f"The name of the {entity}",
    )
