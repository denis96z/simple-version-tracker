from django.db import models


class Project(models.Model):
    id = models.AutoField(verbose_name='ID', primary_key=True)
    name = models.CharField(verbose_name='Name', unique=True, max_length=127)

    def __str__(self):
        return self.name

    class Meta:
        db_table = 'project'
        ordering = ('id',)
