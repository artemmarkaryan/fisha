import uuid
from django.db import models
from django.contrib.postgres.validators import (RangeMaxValueValidator, RangeMinValueValidator)


class Interest(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=512)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)


class User(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    interests = models.ManyToManyField(to="Interest", through="UserInterest")


class UserInterest(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    user_id = models.ForeignKey(to="User", on_delete=models.CASCADE)
    interest_id = models.ForeignKey(to="Interest", on_delete=models.CASCADE)

    rank = models.FloatField(validators=[RangeMaxValueValidator(1), RangeMinValueValidator(-1)])
