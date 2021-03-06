# Trustix
# Copyright (C) 2021 Tweag IO

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

from trustix_nix_reprod.lib.defer import DeferStack
import itertools
import typing


T = typing.TypeVar("T")


def flatten(seq: typing.Iterable[typing.Iterable[T]]) -> typing.Iterator[T]:
    return itertools.chain.from_iterable(seq)


def unique(seq: typing.Iterable[T]) -> typing.Generator[T, None, None]:
    seen: typing.Set[T] = set()
    for i in seq:
        if i not in seen:
            seen.add(i)
            yield i


__all__ = ("DeferStack", "flatten")
