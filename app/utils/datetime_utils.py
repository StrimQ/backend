from datetime import datetime, timezone


def aware_utcnow():
    """
    Returns the current datetime in UTC timezone
    """
    return datetime.now(timezone.utc)
